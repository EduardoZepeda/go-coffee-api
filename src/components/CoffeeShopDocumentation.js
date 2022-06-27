import React from 'react'

import Stack from '@mui/material/Stack';
import Endpoint from './Endpoint'
import Headings from './Headings';

export default function CoffeeShopDocumentation() {

    const coffeeAPIRoutes = [
        {
            "method": "POST", "uri": "/api/v1/cafes/nearest",
            "summary": "Get a list of the ten nearest coffee shops",
            "description": "Get a list of the user ten nearest coffee shops in Guadalajara, ordered by distance. It needs user's latitude and longitude as float numbers",
            "payload": `{"latitude": -103.3668161, "longitude": 20.6708447}`,
            "permissions": null
        },
        {
            "method": "GET", "uri": "/api/v1/cafes/search/{searchTerm}",
            "summary": "Search coffee shops by name or address",
            "description": "Get a list of all coffee shops whose address or name match the search term.",
            "payload": null,
            "permissions": null
        },
        {
            "method": "GET", "uri": "/api/v1/cafes",
            "summary": "Get a list of coffee shops",
            "description": "Get a list of all coffee shop in Guadalajara. Use page and size GET arguments to regulate the number of objects returned.",
            "payload": null,
            "permissions": null
        },
        {
            "method": "POST", "uri": "/api/v1/cafes", "description": "Create a coffee shop object.",
            "summary": "Create a new coffee shop",
            "payload": `{"name": "coffee shop", "location": [-123.123, 123.123], "address": "False st. 123", "rating": 5.0}`,
            "permissions": "staff"
        },
        {
            "method": "GET", "uri": "/api/v1/cafes/{id}", "description": "Get a specific coffee shop object. Id parameter must be an integer.",
            "summary": "Get a new coffee shop by its id",
            "payload": null,
            "permissions": null
        },
        {
            "method": "PUT", "uri": "/api/v1/cafes/{id}", "description": "Create a coffee shop object.",
            "summary": "Update a coffee shop",
            "payload": `{"name": "coffee shop", "location": [-123.123, 123.123], "address": "False st. 123"}`,
            "permissions": "staff"
        },
        {
            "method": "DELETE", "uri": "/api/v1/cafes/{id}", "description": "Delete a coffee shop object.",
            "summary": "Delete a coffee shop",
            "payload": null,
            "permissions": "staff"
        },
    ]

    return (
        <>
            <Headings subtitle={"Coffee shop documentation"} />
            <Stack spacing={2}>
                {coffeeAPIRoutes.map((route) => <Endpoint key={`${route.method}-${route.uri}`} {...route} />)}
            </Stack>
        </>
    )
}
