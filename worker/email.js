export async function resend(env, data) {
    const RESEND_API_KEY = env.RESEND_API_KEY;

    const datas = {
      from: 'Resend <resend@notifications.thesixonenine.site>',
      to: ['thesixonenine@outlook.com'],
      subject: 'Notify from Cloudflare Worker',
    //   html: `<h1>Hello!</h1><p>This email was sent via Cloudflare Worker and Resend API.</p>`,
      html: data,
    };

    const response = await fetch('https://api.resend.com/emails', {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${RESEND_API_KEY}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(datas),
    });
    if (!response.ok) {
      const errorText = await response.text();
      return new Response(`Failed to send email: ${errorText}`, { status: 500 });
    }
    return new Response('success', { status: 200 });
}
