'use client'

// pages/create-token.js
import { useState } from 'react';
import { Button } from '@/components/ui/button';

export default function CreateToken() {
  const [sender, setSender] = useState('');
  const [receiver, setReceiver] = useState('');
  const [amount, setAmount] = useState(0);
  const [tokenType, setTokenType] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    const response = await fetch('/api/create-token', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ sender, receiver, amount, tokenType }),
    });
    const data = await response.json();
    console.log('Token created:', data);
  };

  return (
    <div>
      <h1>Create Token Transaction</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Sender"
          value={sender}
          onChange={(e) => setSender(e.target.value)}
        />
        <input
          type="text"
          placeholder="Receiver"
          value={receiver}
          onChange={(e) => setReceiver(e.target.value)}
        />
        <input
          type="number"
          placeholder="Amount"
          value={amount}
          onChange={(e) => setAmount(Number(e.target.value))}
        />
        <input
          type="text"
          placeholder="Token Type"
          value={tokenType}
          onChange={(e) => setTokenType(e.target.value)}
        />
        <Button type="submit">Create Token</Button>
      </form>
    </div>
  );
}
