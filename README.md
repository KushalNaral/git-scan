# Git Scanner

A command-line tool for scanning Git repositories to analyze commit history and generate statistics based on email addresses.

## Overview

Git Scanner is a simple yet powerful tool that allows you to:
- Scan specific folders containing Git repositories
- Generate commit statistics for a given email address
- Analyze Git history and contribution patterns

## Installation

```bash
go get github.com/yourusername/git-scanner
```

## Usage

The tool provides two main functionalities:

### 1. Adding a New Folder for Scanning

```bash
git-scanner -add /path/to/repository
```

This command will:
- Verify if the specified path contains a Git repository
- Add it to the scanning index
- Process the repository for future statistics generation

### 2. Generating Email Statistics

```bash
git-scanner -email developer@example.com
```

This command will:
- Search through all indexed repositories
- Generate statistics for commits associated with the specified email
- Display a summary of contributions

### Command-line Arguments

| Flag    | Description                     | Default          | Required |
|---------|--------------------------------|------------------|----------|
| `-add`  | Path to Git repository folder  | ""              | No       |
| `-email`| Email address for statistics   | "your@email.com"| No       |

## Examples

Scan a new repository:
```bash
git-scanner -add ~/projects/my-repo
```

Get statistics for a specific email:
```bash
git-scanner -email john.doe@company.com
```

## Implementation Details

The tool consists of two main functions:

1. `scan(folder string)`: 
   - Processes a new Git repository
   - Indexes commit history
   - Stores repository metadata

2. `stats(email string)`:
   - Retrieves commit statistics
   - Analyzes contribution patterns
   - Generates summary reports

## Requirements

- Go 1.x or higher
- Git installed and accessible from command line
- Read access to Git repositories

