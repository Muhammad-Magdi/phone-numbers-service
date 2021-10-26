# Phone Numbers Service

## Requirements

- [X] Validate numbers using RegExp.
- [X] Tests.
- [ ] List and categorize country phone numbers.
- [ ] Filters by country and state.
- [ ] Pagination.

## Endpoints

- `GET: /countries` Returns a list of countries `name` and `code`, This is going to be used by the filtering drop-lists if FE.

## Design Decisions

### Code Architecture
This project heavily inspired by [SOLID principles](https://en.wikipedia.org/wiki/SOLID) and [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). Mainly, to avoid coupling between layers.

![Uncle Bob's Clean Architecture Model](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

However, I preferred to use terms like:
- `Model` over `Entity`.
- `Service` over `Use Case`.
- `Repository` over `Data access`.

1. Model Layer:

    We've 3 Models (`Customer`, `Phone` and `Country`).
    
    These models can **only** depend on each other but shouldn't have any other -direct- dependencies from other layers. However, it can depend on injected dependencies from other layers.



### Database
I have considered the given DB as a mandatory design, So I didn't try to make any update in it (rather than renaming `customer` table to be `customers`).
