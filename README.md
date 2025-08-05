# Simple bank service

## The service I am building is a simple bank. It will provide APIs for the frontend to do following things:
1. Create and manage bank accounts, which are composed of owner’s name, balance, and currency.
2. Record all balance changes to each of the account. So every time some money is added to or subtracted from the account, an account entry record will be created.
3. Perform a money transfer between 2 accounts. This should happen within a transaction, so that either both accounts’ balance are updated successfully or none of them are.

# Tools and Technologies used;
- Golang, Postgres, Redis, Gin, gRPC, Docker, Kubernetes, AWS, CI/CD(GitHub actions), Auth-JWT/Paseto, Unit Tests-gomock

### Credits
Backend Master Class course by TECH SCHOOL
<img width="960" height="540" alt="backend-master" src="https://github.com/user-attachments/assets/12044d5b-e02e-4edc-8e70-fbd5105f0049" />

##### Author
Created by @nfiawornu

###### License
MIT
