import React from 'react'
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import Typography from '@mui/material/Typography';
import Stack from '@mui/material/Stack';
import Endpoint from './Endpoint'

export default function Content() {

    const coffeeAPIRoutes = [
        {
            "method": "GET", "uri": "/api/v1/cafes",
            "description": "Get a list of all coffee shop (cafe) in Guadalajara. Use page and size GET arguments to regulate the number of objects returned.",
            "payload": null,
            "permissions": null
        },
        {
            "method": "POST", "uri": "/api/v1/cafes", "description": "Create a coffee shop (cafe) object.",
            "payload": `{"name": "Cafe", "location": [-123.123, 123.123], "address": "False st. 123", "rating": 5.0}`,
            "permissions": "staff"
        },
        {
            "method": "GET", "uri": "/api/v1/cafes/{id}", "description": "Get a specific coffee shop (cafe) object. Id parameter must be an integer.",
            "payload": null,
            "permissions": null
        },
        {
            "method": "PUT", "uri": "/api/v1/cafes/{id}", "description": "Create a coffee shop (cafe) object.",
            "payload": `{"name": "Cafe", "location": [-123.123, 123.123], "address": "False st. 123"}`,
            "permissions": "staff"
        },
        {
            "method": "DELETE", "uri": "/api/v1/cafes/{id}", "description": "Delete a coffee shop (cafe) object.",
            "payload": null,
            "permissions": "staff"
        },
    ]

    return (
        <Container>
            <Box>
                <Typography variant="h3" component="h1" align="center" paragraph>
                    Coffee API Gdl V1
                </Typography>
                <Stack spacing={2}>
                    {coffeeAPIRoutes.map((route) => <Endpoint {...route} />)}
                </Stack>
            </Box>
        </Container>
    )
}
