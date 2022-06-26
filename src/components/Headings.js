import React from 'react'
import { Typography } from '@mui/material'


export default function Headings({ title, subtitle }) {

    return (
        <>
            <Typography sx={{ my: 2 }} variant="h4" align="center" paragraph>
                {subtitle}
            </Typography>
        </>
    )
}
