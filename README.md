# AI Agent with Claude

A Go application that creates an interactive chat interface with Claude AI (by Anthropic) with the ability to execute tools for enhanced functionality.

## Features

- Interactive chat with Claude 3.7 Sonnet
- Tool execution capability
- Ability to edit files

## Prerequisites

- Go 1.24 or higher
- Anthropic API key

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/kvn-alcantara/ai-agent.git
   cd ai-agent
   ```

2. Copy the example environment file and add your Anthropic API key:
   ```bash
   cp .env.example .env
   ```
   Then edit `.env` to add your Anthropic API key

3. Build and run the application:
   ```bash
   go run main.go
   ```

## Project Structure

- [`main.go`](main.go): Core application with agent implementation
- [`tool/tool_definition.go`](tool/tool_definition.go): Tool interface definition
- [`tool/read_file_tool.go`](tool/read_file_tool.go): File reading tool implementation
- [`tool/list_files_tool.go`](tool/list_files_tool.go): Directory listing tool implementation
- [`tool/edit_file_tool.go`](tool/edit_file_tool.go): File editing tool implementation

## Usage

Once running, you can chat with Claude directly in the terminal. The agent has access to tools like reading and editing files from your system.

Example:
```
You: Create a php script that can define if a string is a palindrome
Claude: I can create a PHP script that checks if a string is a palindrome. A palindrome is a word, phrase, number, or other sequence of characters that reads the same forward and backward (ignoring spaces, punctuation, and capitalization).

Let me create this script for you:
tool: edit_file({"path":"palindrome_checker.php","old_str":""})
Claude: Let me create the PHP script to check for palindromes:
tool: edit_file({"path":"palindrome_checker.php" ...})
Claude: Now let's verify that the file was created properly:
tool: read_file({"path":"palindrome_checker.php"})
Claude: I've created a PHP script called `palindrome_checker.php` that checks if a string is a palindrome. Here's an explanation of how it works:

### Key Features:

1. **The `isPalindrome()` Function**:
   - Converts the string to lowercase
   - Removes all non-alphanumeric characters (spaces, punctuation, etc.)
   - Compares the cleaned string with its reverse
   - Returns `true` if they match, `false` otherwise
...
```

