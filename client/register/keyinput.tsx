import * as React from 'react';
import { Button, Container, Header, Form, Icon } from 'semantic-ui-react';

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
      <Form>
    <Form.Field>
      <label>Key</label>
      <input placeholder='First Name' />
    </Form.Field>
    <Button fluid type='submit'>Submit</Button>
    </Form>
    );
  }
}
