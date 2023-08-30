import NextAuth from 'next-auth';
import GithubProvider from 'next-auth/providers/github';

export default NextAuth({
  // Configure one or more authentication providers
  providers: [
    GithubProvider({
      clientId: process.env.GITHUB_ID,
      clientSecret: process.env.GITHUB_SECRET,
    }),
    // ...add more providers here
  ],
});

//http://localhost:3000/api/auth/providers
//{"github":
// {"id":"github",
// "name":"GitHub",
// "type":"oauth",
// "signinUrl":"http://localhost:3000/api/auth/signin/github",
// "callbackUrl":"http://localhost:3000/api/auth/callback/github"}}

// http://localhost:3000/api/auth/signin/github
// Sign in with GitHub
// →Try signing in with a different account.
// これは[...nextauth].js に clientId と clientSecret に値が設定されていないから。