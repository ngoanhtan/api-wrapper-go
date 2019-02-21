## Create category
Create category for entity for easier management. Category use to group all the same entities into a group (like Folder/ playlist/or tag)
See details [here](https://docs.uiza.io/#create-category).

```golang
var typeCategory = uiza.FolderType
params := &uiza.CategoryCreateParams{
	Name:        uiza.String(""),
	Type:        &typeCategory,
	Description: uiza.String(""),
	Icon:        uiza.String(""),
	OrderNumber:uiza.Int64(1)}

response, _ := category.Create(params)
log.Printf("%s", response)
```

Example Response
```golang
{
    "id": "f932aa79-852a-41f7-9adc-19935034f944",
    "name": "Playlist sample",
    "description": "Playlist description",
    "slug": "playlist-sample",
    "type": "playlist",
    "orderNumber": 3,
    "icon": "https:///example.com/image002.png",
    "status": 1,
    "createdAt": "2018-06-18T04:29:05.000Z",
     "updatedAt": "2018-06-18T04:29:05.000Z"
}
```
## Retrieve category
Get detail of category
See details [here](https://docs.uiza.io/#retrieve-category).

```golang
params := &uiza.CategoryIDParams{ID :uiza.String("Your category ID")}
response, _ := category.Retrieve(params)
log.Printf("%s\n", response)

```

Example Response

```golang
{
    "id": "f932aa79-852a-41f7-9adc-19935034f944",
    "name": "Playlist sample",
    "description": "Playlist description",
    "slug": "playlist-sample",
    "type": "playlist",
    "orderNumber": 3,
    "icon": "https:///example.com/image002.png",
    "status": 1,
    "createdAt": "2018-06-18T04:29:05.000Z",
     "updatedAt": "2018-06-18T04:29:05.000Z"
}
```
## Retrieve category list
Get all category
See details [here](https://docs.uiza.io/#retrieve-category-list).
```golang
response, _ := category.List()
for _, v := range response {
    log.Printf("%v\n", v)
}
```

Example Response

```golang
[
    {
        "id": "f932aa79-852a-41f7-9adc-19935034f944",
        "name": "Playlist sample",
        "description": "Playlist desciption",
        "slug": "playlist-sample",
        "type": "playlist",
        "orderNumber": 3,
        "icon": "/example.com/image002.png",
        "status": 1,
        "createdAt": "2018-06-18T04:29:05.000Z",
        "updatedAt": "2018-06-18T04:29:05.000Z"
    },
    {
        "id": "ab54db88-0c8c-4928-b1be-1e7120ad2c39",
        "name": "Folder sample",
        "description": "Folder's description",
        "slug": "folder-sample",
        "type": "folder",
        "orderNumber": 1,
        "icon": "/example.com/icon.png",
        "status": 1,
        "createdAt": "2018-06-18T03:17:07.000Z",
        "updatedAt": "2018-06-18T03:17:07.000Z"
    }
]
```
## Update category
Update information of category
See details [here](https://docs.uiza.io/#update-category).

```golang
var typeCategory = uiza.FolderType
params := &uiza.CategoryUpdateParams{
	ID: uiza.String("Your category ID"),
	Name: uiza.String(""),
	Type: &typeCategory,
	Description:uiza.String(""),
	Icon:uiza.String(""),
	OrderNumber:uiza.Int64(2)}
response, _ := category.Upddate(params)
log.Printf("%s", response)
```

Example Response

```golang
{
    "id": "f932aa79-852a-41f7-9adc-19935034f944",
    "name": "Playlist sample",
    "description": "Playlist description",
    "slug": "playlist-sample",
    "type": "playlist",
    "orderNumber": 3,
    "icon": "https:///example.com/image002.png",
    "status": 1,
    "createdAt": "2018-06-18T04:29:05.000Z",
     "updatedAt": "2018-06-18T04:29:05.000Z"
}
```
## Delete category
Delete category
See details [here](https://docs.uiza.io/#delete-category).
```golang
params := &uiza.CategoryIDParams{ID: uiza.String("Your category ID")}
response, _ := category.Delete(params)
log.Printf("%s", response)
```

Example Response

```golang
{
   "id": "095778fa-7e42-45cc-8a0e-6118e540b61d"
}
```
## Create category relation
Add relation for entity and category.
See details [here](https://docs.uiza.io/#create-category-relation).
```golang
params := &uiza.CategoryRelationParams{
		EntityId: uiza.String(""),
		MetadataIds: []*string{uiza.String(""), uiza.String("")}}
response, _ := category.CreateRelation(params)
for _, v := range response {
	log.Printf("%v\n", v)
}
```

Example Response

```golang
[
    {
        "id": "5620ed3c-b725-4a9a-8ec1-ecc9df3e5aa6",
        "entityId": "16ab25d3-fd0f-4568-8aa0-0339bbfd674f",
        "metadataId": "095778fa-7e42-45cc-8a0e-6118e540b61d"
    },
    {
        "id": "47209e60-a99f-4c96-99fb-be4f858481b4",
        "entityId": "16ab25d3-fd0f-4568-8aa0-0339bbfd674f",
        "metadataId": "e00586b9-032a-46a3-af71-d275f01b03cf"
    }
]
```
## Delete category relation
Delete relation for entity and category
See details [here](https://docs.uiza.io/#delete-category-relation).
```golang
params := &uiza.CategoryRelationParams{
		EntityId: uiza.String(""),
		MetadataIds: []*string{uiza.String(""), uiza.String("")}}
response, _ := category.DeleteRelation(params)
for _, v := range response {
	log.Printf("%v\n", v)
}
```[
        {
            "id": "5620ed3c-b725-4a9a-8ec1-ecc9df3e5aa6",
            "entityId": "16ab25d3-fd0f-4568-8aa0-0339bbfd674f",
            "metadataId": "095778fa-7e42-45cc-8a0e-6118e540b61d"
        },
    ]

Example Response

```golang
[
    {
        "id": "5620ed3c-b725-4a9a-8ec1-ecc9df3e5aa6",
        "entityId": "16ab25d3-fd0f-4568-8aa0-0339bbfd674f",
        "metadataId": "095778fa-7e42-45cc-8a0e-6118e540b61d"
    },
    {
        "id": "47209e60-a99f-4c96-99fb-be4f858481b4",
        "entityId": "16ab25d3-fd0f-4568-8aa0-0339bbfd674f",
        "metadataId": "e00586b9-032a-46a3-af71-d275f01b03cf"
    }
]
```