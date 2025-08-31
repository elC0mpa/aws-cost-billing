# aws-billing

A terminal-based AWS cost and billing tool built with Golang. It provides a costs comparison between current and previous month but in the same period of time. For example, if today is 15th of the month, it will compare costs from 1st to 15th of the current month with costs from 1st to 15th of the previous month.

## Demo

### Basic usage

![](https://github.com/elC0mpa/aws-cost-billing/blob/main/demo/basic.gif)

## Motivation

As a Cloud Architect, I often need to check AWS costs and billing information. Even though AWS provides this information through the console, I usually executed the same steps to get the summary I needed, and basically this is why I created this tool. Besides saving time, it provides a table with all information you need to compare costs between current and previous month for the same period of time.

## Flags

- `--profile`: Specify the AWS profile to use (default is "").
- `--region`: Specify the AWS region to use (default is "us-east-1").

## Pending features

- [ ] Add monthly trend analysis.
- [ ] Export report to CSV and PDF formats.
- [ ] Distribute the CLI in fedora, ubuntu and macOS repositories.
