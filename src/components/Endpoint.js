import * as React from 'react';
import Paper from '@mui/material/Paper';
import { Typography } from '@mui/material';
import { styled } from '@mui/material/styles';
import Alert from '@mui/material/Alert';
import Code from './Code';
export default function Endpoint({ method, uri, summary, description, payload, permissions }) {

    const Item = styled(Paper)(({ theme }) => ({
        ...theme.typography.body2,
        color: theme.palette.text.secondary,
        padding: '1rem',
    }));

    return (
        <Item key={`${method}-${uri}`} elevation={1} variant="overline" display="block">
            <Typography sx={{ my: 1 }} variant="h5">
                {summary}
            </Typography>
            <Alert sx={{ my: 1 }} severity="info" icon={false}>
                {method} {uri}
            </Alert>
            <Typography sx={{ my: 1 }} paragraph>{description}</Typography>
            {payload
                ? (<>
                    <Typography>Payload example:</Typography>
                    <Code code={payload} language={'javascript'} />
                </>)
                : null}
            {permissions ? <Alert sx={{ my: 1 }} severity="warning">{`Only ${permissions} members allowed`}</Alert> : null}
        </Item>
    );
}
