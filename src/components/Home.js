import React from 'react'

import Typography from '@mui/material/Typography';

export default function Home() {



    return (
        <>
            <Typography variant="h3" component="h1" align="center" paragraph>
                Coffee API Gdl V1
            </Typography>
            <Typography variant="h4" component="h2" align="center" paragraph>
                API Description
            </Typography>
            <Typography sx={{ color: 'text.secondary' }} paragraph>
                This project is not only the documentation frontend. Coffee API's endpoints are <strong>completely functional</strong> and can be accessed at /api/v1, served by a backend writen in Go,
            </Typography>
            <Typography sx={{ color: 'text.secondary' }} paragraph>
                This API returns information about speciality coffee shops in Guadalajara, Mexico. please refer to documentation for further details. It was created as a personal project to practice go web server capabilities.
            </Typography>
        </>
    )
}
