# Phone Numbers Service

## Requirements

- [X] Validate numbers using RegExp.
- [X] Tests.
- [X] List and categorize country phone numbers.
- [X] Filters by country and state.
- [X] Simple frontend.

## Usage

```bash
  sudo docker-compose up --build
```

Then, use [localhost:8080](http://localhost:8080/) to open `Phone Numbers Service` UI.

## Endpoints

- `GET: /countries` Returns a list of countries `name` and `code`, This is going to be used by the filtering drop-lists if FE.
- `GET: /customers` Returns a list of customers with fields (`id`, `name`, `phone`, `country`, `is_valid`). Accepts 2 Optional Query Parameters (`country`, `is_valid`) that're used to filter results.

## Design Decisions

### Code Architecture
This project heavily inspired by [SOLID principles](https://en.wikipedia.org/wiki/SOLID) and [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). Mainly, to avoid coupling between layers.

![Uncle Bob's Clean Architecture Model](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

However, I preferred to use terms like:
- `Model` over `Entity`.
- `Service` over `Use Case`.
- `Repository` over `Data access`.

Model Layer:

  We've 3 Models (`Customer`, `Phone` and `Country`).
  
  These models can **only** depend on each other but shouldn't have any other -direct- dependencies from other layers. However, it can depend on injected dependencies from other layers.

### Database
I have considered the given DB as a mandatory design, So I didn't try to make any update in it (rather than renaming `customer` table to be `customers`).
