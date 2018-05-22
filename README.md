## Kombustion

[![Build Status](https://travis-ci.org/KablamoOSS/kombustion.svg?branch=master)](https://travis-ci.org/KablamoOSS/kombustion)
[![](https://tokei.rs/b1/github/kablamooss/kombustion)](https://github.com/kablamooss/kombustion)

> Extend Cloudformation with plugins

Kombustion uses plugins to preprocess and extend your CloudFormation templates.

In addition to generating templates, Kombustion can also create, update and
delete your Cloudformation stacks.

Kombustion has automatic support for new CloudFormation types as they are
released ([how?](docs/generation.md)).

See the [Quick start](docs/quickstart.md) for more details.

## Getting Started

Kombustion is distibuted as binaries for all major platforms.

#### MacOS

[64-bit](http://downloads.kombustion.io/releases/latest/MacOS/64/kombustion)

### Linux

[32-bit](http://downloads.kombustion.io/releases/latest/Linux/32/kombustion) |
[64-bit](http://downloads.kombustion.io/releases/latest/Linux/64/kombustion)

After downloading for MacOS or Linux, you will need to move the `kombustion`
binary into your `$PATH`, and make it executable.

```bash
sudo chmod +x kombustion
sudo cp kombustion /usr/local/bin/kombustion
```

### Windows

[32-bit](http://downloads.kombustion.io/releases/latest/Windows/32/kombustion.exe)
|
[64-bit](http://downloads.kombustion.io/releases/latest/Windows/64/kombustion.exe)

### Docker

Alternatively, you can run Kombustion via our public Docker image:

```sh
docker run -ti kablamooss/kombustion
```

## Usage

### Cloudformation Stacks

A stack template is written in the same way as standard Cloudformation.
Kombustion allows plugins to extend the syntax, but the end result is always
standard Cloudformation.

The following example shows how a small definition, can be processed into a
bigger template. This lets your plugin maintain safe, sane defaults, and ensure
you don't miss any required fields.

```yaml
# In this example we're going to create a bastion host.
# This is a small EC2 instance, configured with a public IP
# and a security group to allow us to SSH into our AWS cloud.
AWSTemplateFormatVersion: 2010-09-09
Description: Example EC2 Instance
Parameters: {}
Mappings: {}
Resources:
  BastionHost:
    Type: Kombustion::Examples::BastionHost
    Properties:
      # In this example, this key would have been uploaded to AWS
      KeyName: my-ssh-key
      Size: t2.micro
      # Using a filter, find the most recent AMI of Amazon Linux 2
      AmiFilter:
        VirtualizationType: "hvm"
        Name: "amzn2-ami-*",
        RootDeviceType: "ebs"
        owners: ["amazon"],
        Latest: true
```

In this case the plugin `Kombustion::Examples::BastionHost` processing the above
template. Finds the correct AMI to use, and generates the following template.

```yaml
AWSTemplateFormatVersion: 2010-09-09
Description: Example EC2 Instance
Parameters:
  KombustionExampleBastionHostKeyName:
    Description: Name of an existing EC2 KeyPair to enable SSH access to the instances
    Type: 'AWS::EC2::KeyPair::KeyName'
    Default: 'my-ssh-key'
    ConstraintDescription: must be the name of an existing EC2 KeyPair.
  KombustionExampleBastionHostSSHLocation:
    Description: The IP address range that can be used to SSH to the EC2 instances
    Type: String
    MinLength: '9'
    MaxLength: '18'
    Default: 0.0.0.0/0
    AllowedPattern: '(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})/(\d{1,2})'
    ConstraintDescription: must be a valid IP CIDR range of the form x.x.x.x/x.
Mappings: {}
Resources:
Resources:
  EC2Instance:
    Type: 'AWS::EC2::Instance'
    Properties:
      InstanceType: !Ref InstanceType
      SecurityGroups:
        - !Ref InstanceSecurityGroup
      KeyName: !Ref KombustionExampleBastionHostKeyName
      ImageId: 'ami-c267b0a0'
  InstanceSecurityGroup:
    Type: 'AWS::EC2::SecurityGroup'
    Properties:
      GroupDescription: Enable SSH access
      SecurityGroupIngress:
        - IpProtocol: tcp
          FromPort: '22'
          ToPort: '22'
          CidrIp: !Ref KombustionExampleBastionHostSSHLocation
  IPAddress:
    Type: 'AWS::EC2::EIP'
  IPAssoc:
    Type: 'AWS::EC2::EIPAssociation'
    Properties:
      InstanceId: !Ref EC2Instance
      EIP: !Ref IPAddress
```

Check out the [configs](https://github.com/KablamoOSS/Kombustion/configs/)
directory for examples.

### Cloudformation Stack Management

Cloudformation is the only target available, support for others may be added
later. For now to use Cloudformation commands start with `kombustion cf`.

Upsert a CloudFormation template:

```sh
kombustion cf upsert configs/test.yaml --stackName test-stack
```

Delete a CloudFormation stack:

```sh
kombustion cf delete configs/test.yaml
```

Print all the events for a stack:

```sh
kombustion cf events configs/test.yaml
```

#### Credentials

Kombustion uses the same method as the `aws` cli to get
[credential information](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html).
You can either use the standard environment variables `AWS_ACCESS_KEY_ID`,
`AWS_SECRET_ACCESS_KEY`, and `AWS_SESSION_TOKEN`.

Or use a profile you have configured, for example:

```sh
kombustion cf upsert configs/test.yaml --stackName test-stack --profile myAwsProfile
```

### Plugins

> Kombustion plugins are not yet supported on Windows, due to
> [this issue](https://github.com/golang/go/issues/19282). Please use Docker or
> WSL in the meantime.

Install a plugin:

```sh
kombustion cf plugins get mypluginname
```

List all installed plugins:

```sh
kombustion cf plugins list
```

Delete an installed plugin:

```sh
kombustion cf plugins delete mypluginname
```

## Contributing

Please read
[CONTRIBUTING.md](https://github.com/kablamoos/kombustion/blob/master/CONTRIBUTING.md)
for details on our code of conduct, and the process for submitting pull requests
to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available,
see the [tags on this repository](https://github.com/kablamoos/kombustion/tags).

## Maintainers

Kombustion is primarily maintained by the [Kablamo](https://www.kablamo.com.au/)
team. Pull requests are welcome.

## Acknowledgements

The Kombustion logo is based on an original design by Renee French.

## License

This project is licensed under the
[MIT License](https://github.com/kablamoos/kombustion/blob/master/LICENSE).

---

Made with :heart: in Australia.
