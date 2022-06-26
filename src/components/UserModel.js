import React from 'react'

import Stack from '@mui/material/Stack';
import ModelDescription from './ModelDescription'
import Headings from './Headings';

export default function CafeModel() {

    const userModel = [
        { "field": "id", "description": "User's unique Id consist of an integer, autoincremental", "type": "int", "null": "", "blank": "", "limits": "" },
        { "field": "email", "description": "User's email in the form user@provider.com", "type": "string", "null": "", "blank": "", "limits": "" },
        { "field": "password", "description": "User's password, must be longer than 8 characters", "type": "string", "null": "", "blank": "", "limit": "" },
        { "field": "username", "description": "User's username, optional. Currently not used", "type": "string", "null": "", "limit": "" },
    ]


    return (
        <>
            <Headings subtitle={"Coffee shop Model"} />
            <Stack spacing={2}>
                {userModel.map((field) => <ModelDescription key={field.field} field={field} />)}
            </Stack>
        </>
    )
}
