import * as React from 'react';
import CssBaseline from '@mui/material/CssBaseline';
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import { Typography } from '@mui/material';
import { Button } from '@mui/material';

export default function App() {

  return (
    <Box sx={{backgroundColor: 'grey.900', color: 'white', height: '100vh'}}>
      <Container sx={{ display: 'flex', flexDirection:'column',  justifyContent: 'center', alignItems: 'center', height: '100%' }}>
      <CssBaseline />
      <Box sx={{textAlign: 'center'}}>
      <img srcset="cup-360w.png 360w,
             cup-180w.png 180w"
           sizes="(max-width: 480px) 180px,
                  360px"
           src="cup-360w.png" alt="A rainbow cup of coffee" loading="lazy"/>
        <Typography sx={{textAlign:'center'}} variant="h4" component="h1">Go Coffee API</Typography>
        <Typography sx={{textAlign:'center'}}>A social network for coffee lovers</Typography>
      </Box>
      <Box>
      <Button href="/api/v1/swagger/" variant="outlined" size="large" sx={{margin: '24px'}}>
        Go to Coffee API documentation 
      </Button>
      </Box>
      </Container>
    </Box >
  );
}
