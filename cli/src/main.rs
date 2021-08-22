use structopt::StructOpt;

#[derive(Debug, StructOpt)]
struct Opt {
    #[structopt(long)]
    add: bool
}

fn main() {
    let args: Opt = Opt::from_args();

    let _add = args.add;
}