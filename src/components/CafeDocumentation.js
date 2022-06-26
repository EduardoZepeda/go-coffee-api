import React from 'react'

import Typography from '@mui/material/Typography';
import Stack from '@mui/material/Stack';
import Endpoint from './Endpoint'

export default function CafeDocumentation() {

    const coffeeAPIRoutes = [
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
            <Typography variant="h3" component="h1" align="center" paragraph>
                Coffee API Gdl V1
            </Typography>
            <Typography variant="h4" component="h2" align="center" paragraph>
                Cafes documentation
            </Typography>
            <Stack spacing={2}>
                {coffeeAPIRoutes.map((route) => <Endpoint {...route} />)}
            </Stack>
        </>
    )
}
