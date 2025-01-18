# AIHireSphere
> AI-driven Interview Assistance Platform
[中文文档](README_ZH.md)

## Project Introduction
An AI-driven interview assistance platform developed using the [Go-Zero](https://github.com/zeromicro/go-zero) framework.

## Project Technology Stack
1. Go-Zero
2. Gorm
3. MySQL
4. Redis
5. zap
6. ...

## Project Directory Structure
This project follows the DDD (Domain-Driven Design) architecture and is built using the Go-Zero framework as a microservices application.
```
text
.
├── Dockerfile
├── LICENSE
├── README.md
├── README_ZH.md
├── application         // Microservice application layer
│   └── user-center     // User Center
├── common              // Common modules
│   ├── call            // RPC call classes
│   ├── codex           // Response wrapper
│   ├── decorator       // CQRS decorators, not used in this phase
│   ├── gormx           // Gorm operation classes
│   ├── interceptors    // Interceptors
│   ├── model           // Database table classes
│   ├── utils           // Utility classes
│   ├── zapx            // Zap logging classes
│   └── zlog            // Logging service
├── docker-compose.yaml         // One-click environment setup
├── docs                        // Documentation
├── go.mod
├── go.sum
├── logs
├── scripts                 // Script classes
├── template                // Go-Zero templates, to be moved to a public GitHub repository
├── test                    // Test classes
└── utils   
└── path.go
```
For each microservice, our directory structure is as follows:
```
text

.
├── README.md          // Microservice description
├── app                // Application layer
│   └── user.go        // Application
├── domain             // Domain layer
│   ├── events         // Domain events
│   ├── irepository    // Repository interfaces
│   ├── model          // Domain models
│   └── services       // Domain services
├── infrastructure     // Infrastructure layer
│   ├── driver         // Driver layer
│   └── repository     // Repository layer
└── interfaces         // Interface layer
├── api            // API layer
└── rpc            // RPC layer
```
## Before Development

### 1. **Enable Git Hooks**
```
shell
git config core.hooksPath .githooks
chmod -R +x .githooks
```
### 2. **Read the Following Development Guidelines**

#### Branch Naming Convention
We require:

1. Branch names should include the name of the person responsible.
2. Branch names must clearly express the problem being addressed.

Therefore, branch names must be standardized.
```
text
<type>-<name>-<description>
```
For example:
- If it is a branch for developing new features, the naming convention is as follows
```
text
feature-<name>-<feature description>
e.g.: feature-jett-dev_log_system
```
- If it is a bug fix:
```
text
bugfix-<name>-<bug name>
e.g.: bugfix-jett-login_error
```
Other types:
`hotfix`, `release`, ...

#### Commit Message Format
Commit messages should be as clear as possible, with each commit doing one thing.

```
text
<type>(<scope>): <subject>

e.g.: feat: add new api
or: feat(common): add new api
```
#### Types

```
text
# Main Types
feat:      Add new feature
fix:       Fix bug

# Special Types
docs:      Only change documentation-related content
style:     Changes that do not affect the meaning of the code, such as deleting spaces, changing indentation, adding or removing semicolons
build:     Changes to build tools or external dependencies, such as webpack, npm
refactor:  Use when refactoring code
revert:    Message printed when executing git revert

# Types Not Currently Used
test:      Add or modify existing tests
perf:      Improvements to performance
ci:        Changes related to CI (Continuous Integration services)
chore:     Other modifications that do not change src or tests, such as changes to the build process or auxiliary tools
```
#### Subject

No punctuation at the end

For example:
```
text
feat: add new feature
fix: fix a bug
```
## **You Must Know**
1. The project follows the [Uber's Go Style Guide](https://github.com/xxjwxc/uber_go_guide_cn).
2. **Do not** submit any sensitive information in any code, such as `api_key`, `address`, or `password`.
3. You can use the configuration file `config.yaml` to store some sensitive information, but do not try to submit it. Each time you modify the structure of `config.yaml`, you must also update `config.yaml.template`.
4. Do not use `git push --force` unless you know what you are doing.

## Todo List
### Project Phase 1
- [ ] User Center
  - [ ] User Login
  - [ ] User Registration
  - [ ] User Search
- [ ] Interview Center
  - [ ] Resume Upload
  - [ ] Resume Search
  - [ ] Resume Screening
  - [ ] AI Interview (Text Form)
  - [ ] Question Bank Display
  - [ ] Batch Upload Question Bank in Formatted Manner
- [ ] AI Research

### Project Phase 2
- [ ] User Center
  - [ ] User Points
  - [ ] Points Consumption
  - [ ] Points Purchase
- [ ] Interview Center
  - [ ] AI Interview (Voice Form/Video Form)



