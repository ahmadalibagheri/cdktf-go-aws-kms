package main

import "fmt"

func makePolicy(accountId string) string {
	return fmt.Sprintf(
		`{
			"Version": "2012-10-17",
			"Statement": [
			  {
				"Sid": "Enable IAM User Permissions",
				"Effect": "Allow",
				"Principal": {
				  "AWS": "arn:aws:iam::%[1]s:root"
			  },
				"Action": [
				  "kms:*"
				],
				"Resource": [
				  "*"
				]
			  },    {
				"Sid": "Allow autoscalling to use the key",
				"Effect": "Allow",
				"Principal": {
				  "AWS": [
					"arn:aws:iam::%[1]s:role/aws-service-role/autoscaling.amazonaws.com/AWSServiceRoleForAutoScaling"
				  ]
				},
				"Action": [
					"kms:Create*",
					"kms:Describe*",
					"kms:Enable*",
					"kms:List*",
					"kms:Put*",
					"kms:Update*",
					"kms:Revoke*",
					"kms:Disable*",
					"kms:Get*",
					"kms:Delete*",
					"kms:TagResource",
					"kms:UntagResource",
					"kms:ScheduleKeyDeletion",
					"kms:CancelKeyDeletion"
				],
				"Resource": "*"
			  },{
				"Sid": "Allow use of the key",
				"Effect": "Allow",
				"Principal": {
					"AWS": [
					"arn:aws:iam::%[1]s:role/aws-service-role/autoscaling.amazonaws.com/AWSServiceRoleForAutoScaling"
					]
				},
				"Action": [
					"kms:Encrypt",
					"kms:Decrypt",
					"kms:ReEncrypt*",
					"kms:GenerateDataKey*",
					"kms:DescribeKey"
				],
				"Resource": "*"
				},        {
				  "Sid": "Allow attachment of persistent resources",
				  "Effect": "Allow",
				  "Principal": {
					  "AWS": [
				  "arn:aws:iam::%[1]s:role/aws-service-role/autoscaling.amazonaws.com/AWSServiceRoleForAutoScaling"
				  ]
				  },
				  "Action": [
					  "kms:CreateGrant",
					  "kms:ListGrants",
					  "kms:RevokeGrant"
				  ],
				  "Resource": "*",
				  "Condition": {
					  "Bool": {
						  "kms:GrantIsForAWSResource": "true"
					  }
				  }
			  }
			]
		  }`, accountId)
}
