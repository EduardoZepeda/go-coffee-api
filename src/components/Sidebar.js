import React from 'react'
import { styled, useTheme } from '@mui/material/styles';

import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';
import Divider from '@mui/material/Divider';
import Drawer from '@mui/material/Drawer';
import IconButton from '@mui/material/IconButton';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';

import ListItemText from '@mui/material/ListItemText';
import { NavLink } from "react-router-dom";

const upperMenu = [{ "text": "Home", "link": "/" }, { "text": "Coffee shop documentation", "link": "coffee-shop-documentation" }, { "text": "User documentation", "link": "users-documentation" }]
const lowerMenu = [{ "text": "Coffee shop model", "link": "coffee-shop-model" }, { "text": "User model", "link": "user-model" }]


export default function Sidebar({ open, handleDrawerClose, drawerWidth }) {
    const theme = useTheme();

    const DrawerHeader = styled('div')(({ theme }) => ({
        display: 'flex',
        alignItems: 'center',
        padding: theme.spacing(0, 1),
        // necessary for content to be below app bar
        ...theme.mixins.toolbar,
        justifyContent: 'flex-end',
    }));

    return (
        <Drawer
            sx={{
                width: drawerWidth,
                flexShrink: 0,
                '& .MuiDrawer-paper': {
                    backgroundColor: 'text.primary',
                    color: '#FFF',
                    width: drawerWidth,
                    boxSizing: 'border-box',
                },
            }}
            variant="persistent"
            anchor="left"
            open={open}
        >
            <DrawerHeader>
                <IconButton sx={{ color: 'primary.main' }} onClick={handleDrawerClose}>
                    {theme.direction === 'ltr' ? <ChevronLeftIcon /> : <ChevronRightIcon />}
                </IconButton>
            </DrawerHeader>
            <Divider />
            <List>
                {upperMenu.map(({ text, link }) => (
                    <NavLink key={link} style={({ isActive }) =>
                        isActive
                            ? {
                                color: '#29b6f6',
                                textDecoration: 'none'
                            }
                            : { color: 'inherit', textDecoration: 'none' }
                    } exact={'true'} to={link}>
                        <ListItem key={text} disablePadding>
                            <ListItemButton>
                                <ListItemText primary={text} />
                            </ListItemButton>
                        </ListItem>
                    </NavLink>
                ))}
            </List>
            <Divider sx={{ bgcolor: 'primary.dark' }} />
            <List>
                {lowerMenu.map(({ text, link }) => (
                    <NavLink key={link} style={({ isActive }) =>
                        isActive
                            ? {
                                color: '#29b6f6',
                                textDecoration: 'none',
                            }
                            : { color: 'inherit', textDecoration: 'none' }
                    } exact={'true'} to={link}>
                        <ListItem key={text} disablePadding>
                            <ListItemButton>
                                <ListItemText primary={text} />
                            </ListItemButton>
                        </ListItem>
                    </NavLink>
                ))}
            </List>
            <Divider sx={{ bgcolor: "primary.dark" }} />
        </Drawer>
    )
}
