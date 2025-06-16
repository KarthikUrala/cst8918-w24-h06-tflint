CST8918 - DevOps: Infrastructure as Code
Prof: Robert McKenney

Hybrid-H06 TFLint
Background
Building on last week's lab, you will add a linter to your Terraform project. This will help you catch errors and enforce best practices in your Terraform code. We will be using TFLint.

What is a linter?
A linter is a tool that analyses source code to flag programming errors, bugs, stylistic errors, and suspicious constructs. The name is derived from the name of the utility, lint for the C language, which was named after the bits of fluff and dirt that collect in the fibers of clothing. Equivalent tools have been created for many other languages (e.g. ESLint for Javascript/TypeScript).

TFLint is a static analysis tool for Terraform modules. It checks your Terraform code for errors and best practices. This lab will guide you through the installation and configuration of TFLint and show you how to use it to improve your Terraform modules.

Why use TFLint?
Consistency: enforce best practices and coding standards across your Terraform modules.
Quality: catch errors and potential issues before they become problems.
Security: help avoid common security pitfalls in your infrastructure code.
Efficiency: integrate linters into your CI/CD pipeline to automate code quality checks.
Prerequisites
Fork and clone this repository to your local machine. You will need to have Terraform and Go installed. You will also need to have the Azure CLI installed and be logged in to your Azure account.

Tip

Don't forget to run terraform init in the top level project folder to download the AzureRM provider plugin.

Installation
See the official GitHub repo README file for all installation options. The most common are ...

Bash script (Linux):

curl -s https://raw.githubusercontent.com/terraform-linters/tflint/master/install_linux.sh | bash
Homebrew (macOS):

brew install tflint
Chocolatey (Windows):

choco install tflint
Configuration
Default Rule set
TFLint uses the same Hashicorp Configuration Language (HCL) structure as your Terraform modules. Create a .tflint.hcl file with the following content.

plugin "terraform" {
  enabled = true
  preset  = "recommended"
}
This enables the recommended best practice rules. The documentation provides a quick check list of the rules included in the "recommended" preset and links to each rule's more detailed explanation.

To enable a more strict set of rules, change the preset in your .tflint.hcl to all.

Rule set for the AzureRM Provider
Each of the main public cloud providers has a TFLint rule set extension for the unique features of that Terraform provider module. We need the one for the AzureRM Provider. There are over 200 rules in this rule set. Be sure to review the documentation.

Add this block to your .tflint.hcl file.

plugin "azurerm" {
  enabled = true
  version = "0.25.1"
  source  = "github.com/terraform-linters/tflint-ruleset-azurerm"
}
You can now run tflint --init from your project folder in the terminal to install the required rule sets.

Test it out
Once everything is configured, you can run tflint in the terminal to see what it reports. Using the recommended preset, you should see a single warning ...

❯ tflint
1 issue(s) found:

Warning: `region` variable has no type (terraform_typed_variables)

  on variables.tf line 7:
   7: variable "region" {

Reference: https://github.com/terraform-linters/tflint-ruleset-terraform/blob/v0.5.0/docs/rules/terraform_typed_variables.md
The warning message clearly identifies the error type and location. If you need further explanation of the error, follow the Reference link to the ruleset documentation.

In this case, the error is that the region variable in variables.tf does not have a type. This is a best practice to ensure that the variable is used correctly in the module. The fix is to add a type to the variable definition.

variable "region" {
  type    = string
  default = "westus3"
}
Lab Tasks
Create a new git branch called tflint.
Change the Terraform preset from recommended to all.
Correct all reported issues.
Make sure that the Terratest conditions are still passing.
[!IMPORTANT] Don't forget to update the subscriptionID and labelPrefix variables in the test file.

Make a git commit and push your branch to GitHub.
Submit your GitHub repo's URL on Brightspace.
