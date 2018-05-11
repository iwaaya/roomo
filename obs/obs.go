package obs

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type OBS struct {
	svc    *s3.S3
	bucket *string
}

type Config struct {
	AccessKey string `yaml:accesskey`
	SecretKey string `yaml:secretkey`
	Region    string `yaml:region`
	URL       string `yaml:url`
	Bucket    string `yaml:bucket`
}

func New(c Config) (*OBS, error) {
	creds := credentials.NewStaticCredentials(c.AccessKey, c.SecretKey, "")
	cfg := &aws.Config{
		Region:      aws.String(c.Region),
		Credentials: creds,
		Endpoint:    aws.String(c.URL),
	}
	sess, err := session.NewSession(cfg)
	if err != nil {
		return nil, err
	}
	svc := s3.New(sess)
	return &OBS{svc, aws.String(c.Bucket)}, nil
}

func (o *OBS) PutObject(key string, body io.ReadSeeker) error {
	_, err := o.svc.PutObject(&s3.PutObjectInput{
		Bucket: o.bucket,
		Key:    aws.String(key),
		Body:   body,
	})
	return err
}
