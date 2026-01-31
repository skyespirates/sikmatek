import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useMutation } from "@tanstack/react-query";
import {
  generateKey,
  encryptText,
  decryptText,
} from "@/services/encrypt-decrypt";
import { useState } from "react";
const EncryptDecrypt = () => {
  const generateKeyMutation = useMutation({
    mutationFn: generateKey,
    onSuccess: (data) => {
      setKey(data.key);
    },
  });
  const encryptTextMutation = useMutation({
    mutationFn: encryptText,
    onSuccess: (data) => {
      setEncryptedResult(data.encrypted);
    },
  });
  const decryptTextMutation = useMutation({
    mutationFn: decryptText,
    onSuccess: (data) => {
      setDecryptedResult(data.decrypted);
    },
  });

  const [key, setKey] = useState("");
  const [text, setText] = useState("");
  const [encrypted, setEncrypted] = useState("");

  const [encryptedResult, setEncryptedResult] = useState("");
  const [decryptedResult, setDecryptedResult] = useState("");
  return (
    <div className="min-h-dvh flex justify-center items-baseline">
      <div className="flex flex-col gap-4 w-3xl border-2 mt-12 p-4 rounded-sm shadow-sm">
        <div className="border-2 rounded-sm shadow-sm">
          <div className="flex flex-col justify-center gap-4">
            <h1 className="text-center text-4xl">{key}</h1>
            <Button
              className="self-center cursor-pointer"
              onClick={() => {
                generateKeyMutation.mutate();
              }}
            >
              Generate Key
            </Button>
            <div></div>
          </div>
        </div>

        <div className="border-2 p-4 rounded-sm shadow-sm">
          <h1>Encryption</h1>
          <div className="flex gap-1 mb-4">
            <Input placeholder="key" value={key} disabled />
            <Input
              placeholder="text"
              value={text}
              onChange={(e) => setText(e.target.value)}
            />
            <Button
              className="cursor-pointer"
              onClick={() => {
                if (!key || !text.trim()) {
                  return;
                }
                encryptTextMutation.mutate({ key: key, text: text.trim() });
              }}
            >
              Encrypt
            </Button>
          </div>
          <div className="p-4 bg-gray-200">
            <p className="text-center">{encryptedResult}</p>
          </div>
        </div>

        <div className="border-2 p-4 rounded-sm shadow-sm">
          <h1>Decryption</h1>
          <div className="flex gap-1 mb-4">
            <Input placeholder="key" value={key} disabled />
            <Input
              placeholder="encrypted text"
              value={encrypted}
              onChange={(e) => setEncrypted(e.target.value)}
            />
            <Button
              className="cursor-pointer"
              onClick={() => {
                if (!key || !encrypted.trim()) {
                  return;
                }
                decryptTextMutation.mutate({
                  key: key,
                  text: encrypted.trim(),
                });
              }}
            >
              Decrypt
            </Button>
          </div>
          <div className="p-4 bg-gray-200">
            <p className="text-center">{decryptedResult}</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default EncryptDecrypt;
