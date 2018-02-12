import * as React from 'react';
import { BrowserRouter, Route } from 'react-router-dom';
import { Admin, Register } from './pages';

export const App: React.StatelessComponent<{}> = (props) => {
  return (
    <BrowserRouter>
      <div>
        <Route path="/register" component={Register} />
        <Route path="/admin" component={Admin} />
      </div>
    </BrowserRouter>
  );
}
