import * as React from 'react';
import Paper from '@mui/material/Paper';
import { Typography } from '@mui/material';
import { styled } from '@mui/material/styles';
import Alert from '@mui/material/Alert';
import { Grid } from '@mui/material';

export default function Endpoint({ method, uri, description, payload, permissions }) {

    const Item = styled(Paper)(({ theme }) => ({
        ...theme.typography.body2,
        color: theme.palette.text.secondary,
    }));

    return (
        <Item elevation={1} variant="overline" display="block">
            <Typography sx={{ my: 1 }} variant="h4">
                {method}
            </Typography>
            <Alert sx={{ my: 1 }} severity="info" icon={false}>
                {uri}
            </Alert>
            <Typography sx={{ my: 1 }} paragraph>{description}</Typography>
            {payload
                ? (<Grid container direction="row" alignItems="center">
                    <Typography sx={{ my: 1 }}>{`Example payload: ${payload}`}</Typography>
                </Grid>)
                : null}
            <Typography sx={{ my: 1 }}>{permissions ? <Alert severity="warning">{`Only ${permissions} members allowed`}</Alert> : null}</Typography>
        </Item>
    );
}
