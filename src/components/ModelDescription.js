import * as React from 'react';
import Paper from '@mui/material/Paper';
import { styled } from '@mui/material/styles';
import { Typography } from '@mui/material';

export default function Endpoint({ field }) {

    const Item = styled(Paper)(({ theme }) => ({
        ...theme.typography.body2,
        color: theme.palette.text.secondary,
    }));

    const keys = Object.keys(field)

    return (
        <Item elevation={1} variant="overline" display="block">
            {keys.map((key) => {
                if (key === 'field') {
                    return <Typography variant="h5" component="h3">{field[key]}</Typography>
                }
                return <Typography>{field[key] ? key + ": " + field[key] : null}</Typography>
            })}
        </Item>
    );
}
