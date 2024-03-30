# pass-gen: Password Generator CLI

This is a simple command-line interface (CLI) application written in Go for generating passwords. It allows you to customize the length of the password, include symbols, digits, and letters, and generate multiple passwords at once. Additionally, you can choose to output the passwords as plain text or as SHA256 hashes.

## Installation

To use this password generator, you need to have Go >= 1.22.1 installed on your system. You can then install the generator by running:

```bash
go install github.com/guilycst/pass-gen/cmd/passgen
```

## Usage

The password generator can be invoked from the command line with various options:

```bash
pass-gen [options]
```

### Options

- `-l`: Specify the length of the password (default is 20).
- `-b`: Specify the number of passwords to generate (default is 1).
- `-u`: Specify which character types to include in the password:
  - `l`: Include letters.
  - `s`: Include symbols.
  - `d`: Include digits.
  (You can combine these options, for example: `-u lsd` to include letters, symbols, and digits)
- `-n`: Do not output the trailing newline (ignored if batch size is bigger than one).
- `-h`: Output SHA256 hash of the password.

### Examples

Generate a single password with default options:
```bash
pass-gen
```

Generate a password with a length of 15 characters:
```bash
pass-gen -l 15
```

Generate 5 passwords with symbols and digits only:
```bash
pass-gen -b 5 -u sd
```

Generate a password and output its SHA256 hash:
```bash
pass-gen -h
```

## License

This project is licensed under the Apache License Version 2.0 - see the [LICENSE](LICENSE) file for details.
