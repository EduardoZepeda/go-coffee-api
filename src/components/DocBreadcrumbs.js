import * as React from 'react';
import Typography from '@mui/material/Typography';
import Breadcrumbs from '@mui/material/Breadcrumbs';
import { Link } from 'react-router-dom';
import { useLocation } from 'react-router-dom';

export default function DocBreadcrumbs() {
    let location = useLocation();
    return (
        <div role="presentation">
            <Breadcrumbs aria-label="breadcrumb">
                <Link style={{ textDecoration: 'none', color: 'inherit' }} to="/">
                    Home
                </Link>
                <Typography>{location.pathname.slice(1)}</Typography>
            </Breadcrumbs>
        </div>
    );
}