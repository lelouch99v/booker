import Link from 'next/link';

const Index = () => {
  return (
    <div>
      <h1>Booker</h1>
      <p>hello, booker.</p>
      <Link href='/login'>
        <a>ログイン</a>
      </Link>
    </div>
  );
};

export default Index;