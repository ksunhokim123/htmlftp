import * as React from 'react';
import { Button, Container, Header, Icon } from 'semantic-ui-react';

interface State {
  members: number;
}

export class KeyInput extends React.Component<{}, State> {
  public state: any;
  constructor(props: any) {
    super(props);
    this.state = { members: 1 };
  }
  public render() {
    return (
      <div>
      <Container>
      <Button animated='vertical'>
        <Button.Content hidden>Source</Button.Content>
        <Button.Content visible>
          <Icon name='github' />
        </Button.Content>
      </Button>
    </Container>
      </div>
    );
  }
}
