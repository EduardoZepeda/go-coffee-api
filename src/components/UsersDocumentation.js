import React from 'react'

import Typography from '@mui/material/Typography';
import Stack from '@mui/material/Stack';
import Endpoint from './Endpoint'

export default function UsersDocumentation() {

    const usersAPIRoutes = [
        {
            "method": "POST", "uri": "/api/v1/login",
            "summary": "Login to an account",
            "description": "Login to user account, requires an email and a password. When login is successful server returns a JWT authorization token.",
            "payload": `{ "email": "email@emailprovider.com", "password": "hyper-secure-password" }`,
            "permissions": null
        },
        {
            "method": "POST", "uri": "/api/v1/signup", "description": "Register a new user in the system, requires an email and a password. Registrations are closed.",
            "summary": "Create a new account",
            "payload": `{ "email": "email@emailprovider.com", "password": "hyper-secure-password" }`,
            "permissions": null
        },
    ]

    return (
        <>
            <Typography variant="h3" component="h1" align="center" paragraph>
                Coffee API Gdl V1
            </Typography>
            <Typography variant="h4" component="h2" align="center" paragraph>
                Users documentation
            </Typography>
            <Stack spacing={2}>
                {usersAPIRoutes.map((route) => <Endpoint {...route} />)}
            </Stack>
        </>
    )
}
