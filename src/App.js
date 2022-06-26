import * as React from 'react';
import { styled } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import {
  Routes,
  Route,
} from "react-router-dom";
import CafeDocumentation from './components/CafeDocumentation';
import DocBreadcrumbs from './components/DocBreadcrumbs';
import Sidebar from './components/Sidebar';
import Navbar from './components/Navbar';
import UsersDocumentation from './components/UsersDocumentation';
import CoffeeShopModel from './components/CoffeeShopModel';
import UserModel from './components/UserModel';
import Home from './components/Home';

const drawerWidth = 240;

const Main = styled('main', { shouldForwardProp: (prop) => prop !== 'open' })(
  ({ theme, open }) => ({
    flexGrow: 1,
    padding: theme.spacing(3),
    transition: theme.transitions.create('margin', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    marginLeft: `-${drawerWidth}px`,
    ...(open && {
      transition: theme.transitions.create('margin', {
        easing: theme.transitions.easing.easeOut,
        duration: theme.transitions.duration.enteringScreen,
      }),
      marginLeft: 0,
    }),
  }),
);


const DrawerHeader = styled('div')(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  padding: theme.spacing(0, 1),
  // necessary for content to be below app bar
  ...theme.mixins.toolbar,
  justifyContent: 'flex-end',
}));

export default function App() {
  const [open, setOpen] = React.useState(false);

  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  return (
    <Box sx={{ display: 'flex' }}>
      <CssBaseline />
      <Navbar open={open} handleDrawerOpen={handleDrawerOpen} drawerWidth={drawerWidth} />
      <Sidebar open={open} handleDrawerClose={handleDrawerClose} drawerWidth={drawerWidth} />
      <Main open={open}>
        <DrawerHeader />
        <DocBreadcrumbs />
        <Container sx={{ display: 'flex', justifyContent: 'center' }}>
          <Box sx={{ maxWidth: '100%' }}>
            <Routes>
              <Route path="/coffee-shop-documentation" element={<CafeDocumentation />} />
              <Route path="/users-documentation" element={<UsersDocumentation />} />
              <Route path="/coffee-shop-model" element={<CoffeeShopModel />} />
              <Route path="/user-model" element={<UserModel />} />
              <Route path="/" element={<Home />} />
            </Routes>
          </Box>
        </Container>
      </Main>
    </Box >
  );
}
