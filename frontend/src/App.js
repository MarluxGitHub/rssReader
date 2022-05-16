import React from 'react';
import { Admin, Resource, ListGuesser } from 'react-admin';
import simpleRestProvider from 'ra-data-simple-rest';

const dataProvider = simpleRestProvider('http://localhost:8080');

const App = () => (
        <Admin dataProvider={dataProvider}>
            <Resource name="feeds" list={ListGuesser} />
        </Admin>
);

export default App;