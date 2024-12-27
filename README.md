# Skill DSL Configuration Language

[English](#english) | [中文](#chinese)

<a name="english"></a>

## English

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

---

<a name="chinese"></a>

## 中文

### 简介

一个为游戏技能配置设计的领域特定语言（DSL）。目前仍处于原型/玩具项目阶段，尚未经过生产环境测试。

### 项目背景

几年前在开发技能系统时，遇到了复杂的技能设计需求，简单的Excel配置已经无法满足需求。这促使了开发专用DSL的想法，期望通过支持语句、表达式、循环等语言特性来实现更丰富的技能逻辑。

当时受限于知识储备和技术能力，采用了Lua作为配置方案。虽然Lua非常优雅且胜任这项工作，但仍存在一些不完善之处：

1. 配置需要引入核心配置之外的额外参数，增加了配置人员的心智负担
2. 无法实现配置与系统逻辑的完全隔离，配置中可以调用非配置范围内的函数

### 当前开发

随着AI技术的发展，我重新尝试解决这个挑战。令人惊讶的是，基础实现仅用了2-3天，其中90%的代码由AI完成，我主要进行了错误修正和优化工作。AI在这个开发过程中展现的能力令人叹服。

### 优势

DSL可以提供：

- 语言层面的非法函数调用限制
- 配置数据校验
- 可控的参数注入
- 以及更多可能性

### 项目结构

- `cmd/`: 主程序
- `lexer/`: 词法分析器
- `syntax/`: 语法分析器
- `ast/`: 抽象语法树节点定义

### 示例

使用示例请参考 `examples/` 目录

运行示例:
```bash
go run cmd/main.go examples/dsl/test_for.dsl examples/output/
```