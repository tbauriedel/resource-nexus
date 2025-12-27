# Development

For simple development, the [compose.yml](./compose.yml) can be used to create the needed environment.

`podman compose --file compose.yml up --detach`

**PostgreSQL:**  
Uses the latest postres image. Port 5432 will be mapped to localhost.  
- User: `resource-nexus`
- Password: `resource-nexus`
- Database name: `resource-nexus`.