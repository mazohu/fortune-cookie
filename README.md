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