# go-cqrs-template
```
- my-project
    - config
        - confog.yaml
        
    - src
        - api
            - cmd
                root.go
                serve.go
            config
                config.go
            - controller
            - middleware
            - route
            - wiring
        - app
            - coin
            - order
                command.go
                command_handler.go
                event_handler.go
                ports.go
                repository.go
            - providers
                - messaging
                    messaging_inferface.go
                - storage
                    redis_interface.go
            - user
        - domain
            - coin
                entity.go
            - order
                entity.go
                event.go
            - user
                entity.go
        - infra
            - providers
                - messaging
                    event_bus.go
                    rabbit_queue.go
                    redis_queue.go
                - storage
                    redis.go
    
        
    
        
    main.go
```