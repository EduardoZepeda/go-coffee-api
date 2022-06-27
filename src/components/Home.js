import React from 'react'

import Typography from '@mui/material/Typography';
import Headings from './Headings';
import { Link } from '@mui/material';
import { Box } from '@mui/system';

export default function Home() {



    return (
        <>
            <Headings subtitle={"API description"} />
            <Typography sx={{ color: 'text.secondary', maxWidth: '76ch', fontSize: '1.3rem' }} paragraph>
                Coffee API's endpoints are <strong>completely functional</strong> and can be accessed at /api/v1, served by a backend writen in Go.
            </Typography>
            <Typography sx={{ color: 'text.secondary', maxWidth: '76ch', fontSize: '1.3rem' }} paragraph>
                This API returns information about speciality coffee shops in Guadalajara, Mexico. Please refer to documentation in the sidebar panel for further details. It was created as a personal project to practice go web server capabilities.
            </Typography>
            <Typography sx={{ color: 'text.secondary', maxWidth: '76ch', fontSize: '1.3rem' }}>Try <Link href='/api/v1/cafes'>/api/v1/cafes</Link></Typography>
            <Box sx={{ display: 'flex', justifyContent: 'center', my: 4 }} >
                <img alt="logo img" src={"logo192.png"}></img>
            </Box>
        </>
    )
}
