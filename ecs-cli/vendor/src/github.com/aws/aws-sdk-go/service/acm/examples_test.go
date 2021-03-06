// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package acm_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acm"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleACM_DeleteCertificate() {
	svc := acm.New(session.New())

	params := &acm.DeleteCertificateInput{
		CertificateArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DeleteCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleACM_DescribeCertificate() {
	svc := acm.New(session.New())

	params := &acm.DescribeCertificateInput{
		CertificateArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DescribeCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleACM_GetCertificate() {
	svc := acm.New(session.New())

	params := &acm.GetCertificateInput{
		CertificateArn: aws.String("Arn"), // Required
	}
	resp, err := svc.GetCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleACM_ListCertificates() {
	svc := acm.New(session.New())

	params := &acm.ListCertificatesInput{
		CertificateStatuses: []*string{
			aws.String("CertificateStatus"), // Required
			// More values...
		},
		MaxItems:  aws.Int64(1),
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.ListCertificates(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleACM_RequestCertificate() {
	svc := acm.New(session.New())

	params := &acm.RequestCertificateInput{
		DomainName: aws.String("DomainNameString"), // Required
		DomainValidationOptions: []*acm.DomainValidationOption{
			{ // Required
				DomainName:       aws.String("DomainNameString"), // Required
				ValidationDomain: aws.String("DomainNameString"), // Required
			},
			// More values...
		},
		IdempotencyToken: aws.String("IdempotencyToken"),
		SubjectAlternativeNames: []*string{
			aws.String("DomainNameString"), // Required
			// More values...
		},
	}
	resp, err := svc.RequestCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleACM_ResendValidationEmail() {
	svc := acm.New(session.New())

	params := &acm.ResendValidationEmailInput{
		CertificateArn:   aws.String("Arn"),              // Required
		Domain:           aws.String("DomainNameString"), // Required
		ValidationDomain: aws.String("DomainNameString"), // Required
	}
	resp, err := svc.ResendValidationEmail(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
