# 🚀 MyHttp

## Introduction

MyHttp Tool is designed to simplify the process of making HTTP requests and hashing the response. By leveraging the tool's parallel execution capability, you can significantly reduce the overall execution time, making it ideal for scenarios involving multiple requests.

The tool comprises three packages: httpclient, semaphore, and urls. The httpclient package handles the HTTP requests, the semaphore package manages concurrent execution, and the urls package provides URL validation and handling utilities.

## Requirements

To run the MyHttp, you need the following:

- Go programming language (version 1.16 or later) installed on your system.
- Internet connectivity to make HTTP requests.
- Basic knowledge of working with command-line tools.

## Usage

Follow the steps below to use the MyHttp:

1.  Clone the repository to your local machine:

```bash
git clone https://github.com/frkntplglu/myhttp.git
```

2.  Change to the project directory:

```bash
cd myhttp
```

3.  Build the executable by running the following command:

```bash
go build -o myhttp
```

This will generate an executable file named **myhttp** in the current directory.

4.  Run the tool with the desired command-line options. The available options are as follows:

```bash
Usage: ./myhttp -flag [args]

Options:

-args
    The URL to make an HTTP request to (required)

--parallel int
    The maximum number of parallel requests (default 10)
```

The args will be your the target URL(s) to make the HTTP request to.

The --parallel option (optional) allows you to specify the maximum number of parallel requests to execute. If not provided, the default limit is set to 10.

Here's an example command that makes a request to **https://example.com** and **http://adjust.com** with a maximum of 5 parallel requests:

```bash
./myhttp --parallel=5 https://example.com adjust.com
```

5.  Upon execution, the tool will make the HTTP request and print the address of the request along with the MD5 hash of the response. The order in which addresses are printed may vary due to parallel execution. This is example of output :

```bash
http://adjust.com b74a489344a7e90ee3b51a19aced96fa
https://example.com 84238dfc8092e5d9c0dac8ef93371a07
```

## Unit Test

The MyHttp includes comprehensive unit tests to ensure the correctness and reliability of its functionalities. To run the unit tests, execute the following command in the project directory:

```bash
go test ./...
```

---

Feel free to explore and enhance the tool as per your requirements. If you encounter any issues or have suggestions for improvements, please don't hesitate to reach out to the project maintainers. Happy requesting!
