import * as React from 'react';
import { KeyInput } from './keyinput'
import { Grid, Segment } from 'semantic-ui-react'

export const RegisterPage: React.StatelessComponent<{}> = (props) => {
  return (
    <div className="registerpage">
    <Grid columns='equal'>
    <Grid.Row>
      <Grid.Column>
        <KeyInput />
        <div>asdfsaf</div>
      </Grid.Column>
      <Grid.Column>
        <Segment>1</Segment>
        <Segment>2</Segment>
        </Grid.Column>
        <Grid.Column>
          <Segment>1</Segment>
          <Segment>2</Segment>
          </Grid.Column>
      </Grid.Row>
      </Grid>
    </div>
  );
}
