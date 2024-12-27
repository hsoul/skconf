# Skill DSL Configuration Language

[English](README.md) | [中文](README_zh_CN.md)

### Introduction

A domain-specific language (DSL) designed for skill configuration in game development. This is currently a prototype/toy project and has not been tested in production environments.

### Background

Several years ago, while developing a skill system, we encountered complex skill design requirements that exceeded the capabilities of simple Excel configurations. This experience inspired the idea of developing a dedicated DSL for skill logic configuration, supporting features like statements, expressions, and loops to enable richer skill logic.

At that time, due to knowledge and technical limitations, we opted for Lua as the configuration solution. While Lua is elegant and capable, it had some imperfections:

1. Configuration required additional parameters beyond the core skill settings, creating extra cognitive load for configuration staff
2. Unable to achieve complete isolation between configuration and system logic, allowing calls to functions outside the configuration scope

### Current Development

With the advancement of AI technology, I revisited this challenge. Surprisingly, the basic implementation took only 2-3 days, with AI contributing to 90% of the code. I mainly focused on error corrections and optimizations. The capabilities of AI in this development process were truly remarkable.

### Benefits

DSL can provide:

- Language-level restrictions on illegal function calls
- Configuration data validation
- Controlled parameter injection
- And many more possibilities

### Project Structure

- `cmd/`: Main program
- `lexer/`: Lexical analyzer
- `syntax/`: Syntax parser
- `ast/`: Abstract Syntax Tree node definitions

### Examples

Please refer to the `examples/` directory for usage examples.

To run the examples:
```bash
go run cmd/main.go examples/dsl/test_for.dsl examples/output/
```

