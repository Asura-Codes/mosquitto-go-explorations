# Gemini MQTT: Hands-on Mosquitto & Go Lab Series

This repository is an experimental project demonstrating the capabilities and potential of the **Eclipse Mosquitto** MQTT broker. All labs and documentation within this repository were autonomously created and configured using the **Gemini Command-Line Interface**, showcasing an AI-driven approach to building comprehensive technical learning environments.

## Overview

The goal of this project is to provide a structured, hands-on curriculum for mastering MQTT concepts using Mosquitto and the Go programming language (via the Paho MQTT client). The labs progress from basic messaging to advanced topics like High Availability, TLS security, and external authentication integration.

### Core Technologies
- **MQTT Broker:** [Eclipse Mosquitto](https://mosquitto.org/)
- **Language:** [Go](https://go.dev/) (using [Paho MQTT](https://github.com/eclipse/paho.mqtt.golang))
- **Infrastructure:** [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)
- **Automation:** PowerShell (`run.ps1`) and VS Code Tasks

## Learning Path

Each lab is self-contained and includes its own configuration, implementation code, and a summary of key takeaways.

| Lab | Topic | Summary |
| :--- | :--- | :--- |
| **Lab 1** | Setting Up the Environment | [View Summary](./lab1/summary.md) |
| **Lab 2** | Go Client - Basic MQTT Messaging | [View Summary](./lab2/summary.md) |
| **Lab 3** | Reliability and Quality of Service (QoS) | [View Summary](./lab3/summary.md) |
| **Lab 4** | Session Persistence and Offline Buffering | [View Summary](./lab4/summary.md) |
| **Lab 5** | Advanced MQTT Patterns (LWT, Req-Rep) | [View Summary](./lab5/summary.md) |
| **Lab 6** | Authentication and Authorization (ACLs) | [View Summary](./lab6/summary.md) |
| **Lab 7** | MQTT v5 Deep Dive | [View Summary](./lab7/summary.md) |
| **Lab 8** | Broker Management and Monitoring ($SYS) | [View Summary](./lab8/summary.md) |
| **Lab 9** | Mosquitto Topologies - Bridging | [View Summary](./lab9/summary.md) |
| **Lab 10** | Application Patterns - State Management | [View Summary](./lab10/summary.md) |
| **Lab 11** | WebSockets and Browser Clients | [View Summary](./lab11/summary.md) |
| **Lab 12** | Advanced Topologies - Aggregation | [View Summary](./lab12/summary.md) |
| **Lab 13** | Security - TLS/SSL Encryption | [View Summary](./lab13/summary.md) |
| **Lab 14** | High Availability - Fan-out Architecture | [View Summary](./lab14/summary.md) |
| **Lab 15** | High Availability - Load Balancing (HAProxy) | [View Summary](./lab15/summary.md) |
| **Lab 16** | High Availability - Server-Side Replication | [View Summary](./lab16/summary.md) |
| **Lab 17** | Dynamic Security Management | [View Summary](./lab17/summary.md) |
| **Lab 18** | External Authentication Integration (Webhooks) | [View Summary](./lab18/summary.md) |

## Prerequisites

To run these labs locally, you will need:
- **Go:** Version 1.18 or later.
- **Docker & Docker Compose:** For running the Mosquitto broker and auxiliary services.
- **PowerShell:** For executing the automation scripts (`run.ps1`).

## Philosophy

- **Clean Code:** Labs adhere to SOLID principles and idiomatic Go patterns.
- **Instructional Focus:** Code is heavily commented to explain the *why* behind MQTT patterns.
- **Production Workflow:** Emphasizes a "Build then Run" workflow to simulate real-world artifact generation.

---
*Constructed using the Gemini CLI.*
