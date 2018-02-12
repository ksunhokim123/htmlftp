import * as React from 'react';
import { KeyInput } from './keyinput';
import { Step, Button, Icon, Grid, Segment, Header } from 'semantic-ui-react';
import { RegisterStep } from './step'

export const RegisterPage: React.StatelessComponent<{}> = (props) => {
  return (
    <div className="registerpage">
      <Header as='h5' attached='top'>
        mouse-ftp register
      </Header>
        <Segment attached>
          <KeyInput />
        </Segment>
        <Segment attached>
        <Button animated='vertical'>
          <Button.Content hidden>Source</Button.Content>
          <Button.Content visible>
            <Icon name='github' />
          </Button.Content>
        </Button>
        </Segment>
    </div>
  );
}
