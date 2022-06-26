import React from 'react'

import Stack from '@mui/material/Stack';
import Endpoint from './Endpoint'
import Headings from './Headings';

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
            <Headings subtitle={"Users documentation"} />
            <Stack spacing={2}>
                {usersAPIRoutes.map((route) => <Endpoint key={`${route.method}-${route.uri}`} {...route} />)}
            </Stack>
        </>
    )
}
