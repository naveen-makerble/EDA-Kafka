# Event-Driven Architecture with Kafka (still working on it)

This project is a small prototype demonstrating the principles of **Event-Driven Architecture (EDA)** using Kafka as the messaging backbone. The system includes a producer for publishing events, a consumer for processing them, and handlers to perform specific tasks.

## How It Works

1. **Producer**:
   - Produces comments (as events) to the Kafka topic named `comments`.

2. **Consumer**:
   - Subscribes to the `comments` topic and processes messages.
   - Uses a handler (`LogHandler`) to log the events to the console.