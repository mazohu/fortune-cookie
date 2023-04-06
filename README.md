# Database Werk :nail_care:
To run the server, run the following commands from `fortune-cookie` folder
```
cd back-end
go run .
```
### Model Tags
When defining `json` and `gorm` metadata on struct members, the metadata is written in order like this:
```golang
type ConfigurationDescription struct {
    ID         int                 `json:"configurationID"`
    Name       string              `json:"name"`
    IsActive   bool                `json:"isActive"`
    LocationID int                 `json:"-"`
    Location   LocationDescription `json:"location,omitempty" gorm:"foreignKey:LocationID;references:ID"`
}
``` 
For many-to-many associations, the join-table owns the foreign keys referring to the schemas. The foreign keys and references are defined with the following tags:
`foreignKey`
: The join-table's foreign key referring to the current schema (i.e. the struct in which this metadata is defined)
`joinForeignKey`
: The join-table's foreign key referring to the associated schema (i.e. the struct defining the attribute on which this metadata is defined)
`References`
: The primary key of the current schema linked by the join-table's `foreignKey`
`joinReferences`
: The primary key of the associated schema linked by the join-table's `joinForeignKey`
## Regarding Golang time package
- "Time values should not be used as map or database keys without first guaranteeing that the identical Location has been set for all values, which can be achieved through use of the UTC or Local method"
- Also worth noting that "local time" is based on the server's system clock, not the client's, so all database times are recorded in reference to the server's timezone