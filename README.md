# AI Agent with Claude

A Go application that creates an interactive chat interface with Claude AI (by Anthropic) with the ability to execute tools for enhanced functionality.

## Features

- Interactive chat with Claude 3.7 Sonnet
- Tool execution capability
- File reading functionality
- Environment variable management

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

## Usage

Once running, you can chat with Claude directly in the terminal. The agent has access to tools like reading files from your system.

Example:
```
echo 'Is AI going to take over the world?' >> theory.txt
```
```
You: What you think about the contents in theory.txt? be brief.
Claude: I'll check the contents of theory.txt and give you my brief thoughts on it.
tool: read_file({"path":"theory.txt"})
Claude: The file "theory.txt" contains a single question: "Is ai going to take over the world?"

In my brief opinion, this question reflects a common concern about artificial intelligence, but the reality is more nuanced than sci-fi scenarios suggest. AI systems are tools designed for specific tasks and lack the general intelligence, consciousness, or motivation needed for "taking over." The real challenges with AI involve ensuring proper governance, addressing bias, and managing economic transitions - not preventing a hostile takeover. The focus should be on developing AI ethically and responsibly to benefit humanity.
```

