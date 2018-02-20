import * as React from 'react';
import { Button, Grid, Header, Icon, Segment, Step } from 'semantic-ui-react';

import { KeyInput } from './keyinput';
import { UserForm } from './userform';

interface State {
  step: number;
  key: string;
  endpoint: string;
}

export class RegisterPage extends React.Component<{}, State> {
  constructor(props: any) {
    super(props);
    this.state = {
      step: 1,
      key: '',
      endpoint: '',
    };
  }

  public userStep(key2: string) {
    this.setState(
      {
        key: key2,
        step: 2,
      }
    );
  }

  public resultStep(endpoint2: string) {
    this.setState(
      {
        endpoint: endpoint2,
        step: 3,
      }
    );
  }

  public render() {
    return (
      <div className='registerpage'>
        <Header as='h5' attached='top'>
          mouse-ftp register
        </Header>
        <Segment attached>
          {
            (() => {
              switch (this.state.step) {
                case 1: {
                  return <KeyInput next={this.userStep.bind(this)}/>;
                }
                case 2: {
                  return <UserForm akey={this.state.key} next={this.resultStep.bind(this)}/>;
                }
                default: {
                  return <div>asdasdads</div>;
                }
              }
            })()
          }
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

};
