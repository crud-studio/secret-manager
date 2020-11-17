package property

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"secret-manager/util"
	"sort"
)

var tagToSearch = secretsmanager.Tag{
	Key:   aws.String("io.doxin.secrets/type"),
	Value: aws.String("applicationProperties"),
}
type SecretValue struct {
	ApplicationProperties string `json:"application.properties"`
}

type Client struct {
	Config *aws.Config
}

func NewClient(config *aws.Config) Client {
	return Client{
		config,
	}
}

func (c Client) GetProperties(name string) (string, error) {
	if !c.isSecretValid(name) {
		return "", errors.New("secret is not valid")
	}

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(name),
	}
	svc := c.getService()
	result, err := svc.GetSecretValue(input)
	if err != nil {
		return "", err
	}

	var secretValue SecretValue

	err = json.Unmarshal([]byte(*result.SecretString), &secretValue)
	if err != nil {
		return "", err
	}

	return secretValue.ApplicationProperties, nil
}

//func (c Client) CreateProperties(name, value string) error {
//	if !c.isSecretValid(name) {
//		return errors.New("secret is not valid")
//	}
//		svc := c.getService()
//
//		secretString, err := json.Marshal(&SecretValue{
//		ApplicationProperties: value,
//	})
//		if err != nil {
//		return err
//	}
//
//	var tags = &[]secretsmanager.Tag{}
//
//	tags = append(*tags, tagToSearch)
//
//	input := secretsmanager.CreateSecretInput{
//		Name: aws.String(name),
//		SecretString: aws.String(string(secretString)),
//		Tags: tags,
//	}
//
//	err := svc.CreateSecret(&input)
//
//		return nil
//	}

func (c Client) ListProperties() ([]string, error) {
	svc := c.getService()

	result, err := svc.ListSecrets(&secretsmanager.ListSecretsInput{})
	if err != nil {
		return nil, err
	}

	var s []string

	for i := range result.SecretList {
		if isTagged(result.SecretList[i].Tags) {
			s = append(s, *result.SecretList[i].Name)
		}
	}

	return s, nil
}

func (c Client) SaveProperties(name, value string) error {
	if !c.isSecretValid(name) {
		return errors.New("secret is not valid")
	}
	svc := c.getService()

	secretString, err := json.Marshal(&SecretValue{
		ApplicationProperties: value,
	})
	if err != nil {
		return err
	}

	input := secretsmanager.PutSecretValueInput{
		SecretId: aws.String(name),

		SecretString: aws.String(string(secretString)),
	}

	svc.PutSecretValue(&input)

	return nil
}

func (c Client) getService() *secretsmanager.SecretsManager {
	s := session.Must(session.NewSession(c.Config))
	svc := secretsmanager.New(s)
	return svc
}

func (c Client) isSecretValid(name string) bool {
	svc := c.getService()
	input := secretsmanager.DescribeSecretInput{
		SecretId: aws.String(name),
	}
	result, err := svc.DescribeSecret(&input)
	if err != nil {
		return false
	}
	return isTagged(result.Tags)
}

func isTagged(tags []*secretsmanager.Tag) bool {
	i := sort.Search(len(tags), func(i int) bool {
		return util.CompareTags(tagToSearch, *tags[i])
	})

	return i != len(tags)
}