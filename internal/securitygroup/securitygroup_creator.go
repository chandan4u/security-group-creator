package securitygroup

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"gopkg.in/yaml.v2"
	"security-group-creator/library"
)

// SGCreator it creates the request security group for there VpcId
func SGCreator(file string) {

	sgYaml, err := sgReader(file)
	if err != nil {
		fmt.Println(fmt.Sprintf("[ERROR] Have some error in reading sg yaml : %v", err))
	}
	ec2client := library.Ec2Session()
	for _, value := range sgYaml.SecurityGroups {
		fmt.Printf("Start process for creating security group name : %s ...\n", value.Name)
		//Create the input struct with the appropriate settings, making sure to use the aws string pointer type
		sgReq := ec2.CreateSecurityGroupInput{
			GroupName:   aws.String(value.Name),
			Description: aws.String(value.Desc),
			VpcId:       aws.String(value.VpcID),
		}

		//Attempt to create the security group
		sgResp, err := ec2client.CreateSecurityGroup(&sgReq)
		if err != nil {
			fmt.Println(fmt.Sprintf("[ERROR] Have some error in creating security group : %v", err))
			continue
		}

		authReq := ec2.AuthorizeSecurityGroupIngressInput{
			CidrIp:     aws.String("0.0.0.0/0"),
			FromPort:   aws.Int64(9443),
			ToPort:     aws.Int64(9443),
			IpProtocol: aws.String("tcp"),
			GroupId:    sgResp.GroupId,
		}
		_, err = ec2client.AuthorizeSecurityGroupIngress(&authReq)
		if err != nil {
			fmt.Println(fmt.Sprintf("[ERROR] Have some error in authorize security group ingress : %v", err))
			continue
		}
		resultOutput(value.Name, *sgResp.GroupId, value.Desc, value.VpcID)
	}
}

// sgReader it basically read the yaml and convert details in struct
func sgReader(file string) (SecurityGroupsCreator, error) {
	var sg SecurityGroupsCreator

	// ReadFromFile : Read the plan text and return value in bytes
	if sgFile, err := library.ReadFromFile(file); err != nil {
		return sg, fmt.Errorf("[ERROR] Have some error with request file : %v", err.Error())
	} else {
		err := yaml.Unmarshal(sgFile, &sg)
		if err != nil {
			return sg, fmt.Errorf("[ERROR] Have some error in unmarshal release yaml : %v", err.Error())
		}
		return sg, nil
	}
}

func resultOutput(sgName, sgId, desc, vpcId string) {
	tw := table.NewWriter()
	headerTransformer := text.Transformer(func(val interface{}) string {
		return text.Bold.Sprint(val)
	})
	tw.AppendHeader(table.Row{"SG Name", "SG Id", "Description", "VpcId"})
	tw.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:              "#",
			WidthMin:          2,
			WidthMax:          4,
			TransformerHeader: headerTransformer,
		},
		{
			Name:              "Account Id",
			WidthMax:          12,
			WidthMin:          12,
			AutoMerge:         true,
			TransformerHeader: headerTransformer,
		},
		{
			Name:              "Cluster",
			WidthMax:          40,
			WidthMin:          40,
			TransformerHeader: headerTransformer,
		},
		{
			Name:              "Description",
			WidthMax:          40,
			WidthMin:          40,
			AutoMerge:         true,
			TransformerHeader: headerTransformer,
		},
	})
	tables := table.Row{
		sgName, sgId, desc, vpcId,
	}
	tw.SetStyle(table.StyleLight)
	tw.Style().Options.SeparateRows = true
	tw.AppendRow(tables)
	tw.SetIndexColumn(1)
	tw.SetTitle("Security Group")
	fmt.Println(tw.Render())
}
