import React from 'react'

import Stack from '@mui/material/Stack';
import ModelDescription from './ModelDescription'
import Headings from './Headings';

export default function CoffeeShopModel() {

    const coffeeShopModel = [
        { "field": "id", "description": "Coffee ship's unique Id consist of an integer, autoincremental", "type": "int", "null": "", "blank": "", "max length": "" },
        { "field": "name", "description": "Coffee shop's name", "type": "string", "null": "", "blank": "", "max length": 100 },
        { "field": "location", "description": "The coordinates of the coffee shop in the form of tuple: [-123.123456, 123.123456]", "type": "Point", "null": "", "blank": "", "max length": "" },
        { "field": "address", "description": "Coffee shop's address, number included", "type": "string", "null": true, "blank": true, "max length": "50" },
        { "field": "rating", "description": "Our main barista rating for that coffee shop. Min 0, Max 5", "type": "Float", "null": true, "blank": true, "max length": "" },
        { "field": "created_date", "description": "Date when coffee shop was registered", "type": "Datetime", "null": "", "blank": "", "max length": "" },
        { "field": "modified_date", "description": "Date when coffee shop was last updated", "type": "Datetime", "null": "", "blank": "", "max length": "" }
    ]


    return (
        <>
            <Headings subtitle={"Coffee shop model"} />
            <Stack spacing={2}>
                {coffeeShopModel.map((field) => <ModelDescription key={field.field} field={field} />)}
            </Stack>
        </>
    )
}
