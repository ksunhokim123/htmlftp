import * as React from 'react';
import { Button, Container, Form, Header, Icon } from 'semantic-ui-react';

interface State {
  key: string;
}

interface Props {
  next(key: string): void;
}

export class KeyInput extends React.Component<Props, State> {
  constructor(props: any) {
    super(props);
    this.state = {
      key: '',
    };
  }

  public onKeyChange(e: any) {
    this.setState(
      {
        key: e.target.value,
      }
    );
  }

  public onSumbit(e: any) {
    e.preventDefault();
    fetch('http://127.0.0.1:5353/api/keys/' + this.state.key, {
          method: 'GET',
          headers: {
            'content-type': 'application/json',
          },
     })
     .then((response) => (response.status))
     .then((status) => {
       if (status === 200) {
         this.props.next(this.state.key);
       }
     });

  }

  public render() {
    return (
      <Form onSubmit={this.onSumbit.bind(this)}>
        <Form.Field>
          <label>Key</label>
          <input placeholder='key' onChange={this.onKeyChange.bind(this)} />
        </Form.Field>
        <Button fluid type='submit' onClick={this.onSumbit.bind(this)}>Submit</Button>
      </Form>
    );
  }
}
