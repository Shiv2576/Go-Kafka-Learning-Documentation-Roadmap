# Stream processing :


## This is a tutorial for stream processing in go:

1. Filtering: only consume messages with certain fields.

2. Transformation: enrich a message before writing to another topic.

3. Fan-out: one topic feeding multiple downstream consumers.


# We are going to implement a real world use case Online Banking Trasaction processing

### A single consumer reads from transactions-raw, then:

1. Filters: Ignores test transactions or internal transfers.

2. Transforms/Enriches:
    1. Adds timestamp: "processed_at": "2024-06-01T10:30:00Z"

    2. Adds business context: "category": "shopping" (based on merchant)

    3. Flags high-risk: "is_high_risk": true (if risk_score > 0.8)

3. Fan-out: Sends this enriched message to three different topics:
      1. transactions.enriched → for data warehouse ingestion (e.g., loaded into Snowflake for analytics)

      2. alerts.fraud → for fraud detection team (only high-risk transactions go here)

      3. user.activity → for customer notification service (e.g., sends “You spent $250 at Amazon” push notification)
