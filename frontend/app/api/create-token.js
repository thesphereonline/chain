import { createTokenTransaction } from '../../lib/blockchain';

export default async function handler(req, res) {
  if (req.method === 'POST') {
    const { sender, receiver, amount, tokenType } = req.body;
    const tx = createTokenTransaction(sender, receiver, tokenType, amount);
    res.status(200).json({ success: true, transaction: tx });
  } else {
    res.status(405).end(); // Method Not Allowed
  }
}
