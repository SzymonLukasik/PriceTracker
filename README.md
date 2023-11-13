# PriceTracker

## Overview
PriceTracker is a web application designed for monitoring product prices across several selected stores. This application allows logged-in users to manage their tracked items through a user-friendly panel.

## Features
- **Product Collection Management**: Users can add products to their personal collection.
- **Price Trend Visualization**: The application provides a graph illustrating the price changes of selected products over time.

## Microservices Architecture
PriceTracker is composed of five microservices:
1. **Scraper**: Responsible for periodically fetching prices of products tracked by at least one user.
2. **Products**: Manages data related to individual product models.
3. **Users**: Stores user preferences and other related data.
4. **Frontend**: User GUI, enabling user interaction with the application and visualization of relevant data. It includes at least two instances, managed by a load balancer.
5. **DiagramGenerator**: Generates price trend graphs for products. It includes an HTTPCache to avoid unnecessary re-rendering of graphics.

## Technologies
- **Frontend**: Golang, gin-gonic
- **Diagram Generator**: Golang, go-echarts
- **Scraper**: Python
- **Database**: PostgreSQL
- **Products, Users**: gRPC

## Communication
- **HTTP**: Frontend ↔ DiagramGenerator
- **gRPC**: Frontend ↔ Users, DiagramGenerator ↔ Products, Scraper ↔ Products

## Architecture Details
- Two instances of the database.
- Three instances of the Frontend.
- One instance each of Scraper, Users, Products, DiagramGenerator.

## Optimizations
- Load balancing before connecting to the Frontend.
- Scalable database with manual sharding of the table containing product price information, sharded by store name.
- Cache: Shared HTTP cache - the diagram generator sends the Frontend a generated diagram with appropriate Cache-Control headers.

## Cache-Control
The Diagram Generator responds to GET requests in the form `/product?shop=[shop]&name=[name]&url=[url]` by sending a generated HTML snippet, including a JS script with the following Cache-Control headers:
- max-age=20
- proxy-revalidate
- stale-while-revalidate
- stale-if-error
- public
- immutable

The response remains fresh for 20 seconds and is shared among users observing the same product.
